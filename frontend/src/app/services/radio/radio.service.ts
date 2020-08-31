import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class RadioService {
  private readonly eventConsumers: {[event: string]: any[]} = {};

  public on(event: string, consumer: any): void {
    if (this.eventConsumers.hasOwnProperty(event)) {
      this.eventConsumers[event].push(consumer);
    } else {
      this.eventConsumers[event] = [consumer];
    }
  }

  public emit<T>(event: string, ...args: T[]): void {
    if (this.eventConsumers.hasOwnProperty(event)) {
      for (const consumer of this.eventConsumers[event]) {
        consumer(...args)
      }
    }
  }
}
