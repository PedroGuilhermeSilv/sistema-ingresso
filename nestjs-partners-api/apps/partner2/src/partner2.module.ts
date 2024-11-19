import { PrismaModule } from '@app/core/prisma/prisma.module';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { EventosModule } from './eventos/eventos.module';
import { LugaresModule } from './lugares/spots.module';
;

@Module({
  imports: [ConfigModule.forRoot({envFilePath: '.env.partner2',isGlobal: true}),
     EventosModule,
      PrismaModule,
       LugaresModule],
})
export class Partner2Module {}
