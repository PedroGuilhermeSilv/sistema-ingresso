import { PrismaService } from '@app/core/prisma/prisma.service';
import { HttpCode, Injectable } from '@nestjs/common';
import { Prisma, TicketStatus } from '@prisma/client';
import { CreateEventDto } from './dto/create-event.dto';
import { ReserveSpotDto } from './dto/reserve-spot.dto';
import { UpdateEventDto } from './dto/update-event.dto';

@Injectable()
export class EventsService {

  constructor(private prismaService: PrismaService){

  }
  create(createEventDto: CreateEventDto) {
    return this.prismaService.event.create({
      data: {
        ...createEventDto,
        date: new Date(createEventDto.date),
      }
    });
  }

  findAll() {
    return this.prismaService.event.findMany();
  }

  findOne(id: string) {
    return this.prismaService.event.findUnique({
      where: {
        id,
      }});
  }

  update(id: string, updateEventDto: UpdateEventDto) {
    return this.prismaService.event.update({
      where: {
        id,
      },
      data: {...updateEventDto,date: new Date(updateEventDto.date)},
    });
  }

  @HttpCode(204)
  remove(id: string) {
    return this.prismaService.event.delete({
      where: {
        id,
        },
        });
  }

  async reserveSpot(dto: ReserveSpotDto & {eventId: string}) {
    const spots = await this.prismaService.spot.findMany({
      where:{
        name: {
          in: dto.spots,
        }
      }
  });
  if (spots.length !== dto.spots.length) {
    const foundSpotsName = spots.map((spot) => spot.name);
    const notFoundSpotsName = dto.spots.filter((spot) => !foundSpotsName.includes(spot));
    throw new Error(`Spots not found: ${notFoundSpotsName.join(', ')}`);
  }
  try {

    const tickets = await this.prismaService.$transaction(async (prisma) => {
      await prisma.reservationHistory.createMany({
        data: spots.map((spot) => ({
          spotId: spot.id,
          ticketKind: dto.ticket_kind,
          email: dto.email,
          status: TicketStatus.reserved,
        })),
      });

      await prisma.spot.updateMany({
        where: {
          id: {
            in: spots.map((spot) => spot.id),
          },
        },
        data: {
          status: TicketStatus.reserved,
        },
      });

      const tickets = await Promise.all(
        spots.map((spot) =>
          prisma.ticket.create({
            data: {
              spotId: spot.id,
              ticketKind: dto.ticket_kind,
              email: dto.email,
            },
          })
        )
      );

      return tickets;
    },{isolationLevel: Prisma.TransactionIsolationLevel.ReadCommitted});
    return tickets;
  }
  catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      switch (error.code) {
        case 'P2002':
        case 'P2034':
          throw new Error('Some of the spots are already reserved')
        default:
          throw error;
      }
    }
  }
  }
}
