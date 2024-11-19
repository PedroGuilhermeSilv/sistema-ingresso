import { SpotsCoreModule } from '@app/core';
import { Module } from '@nestjs/common';
import { LugaresController } from './lugres.controller';


@Module({
  imports: [SpotsCoreModule],
  controllers: [LugaresController],
})
export class LugaresModule {}
