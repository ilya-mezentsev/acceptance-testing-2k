import { Injectable } from '@angular/core';
import {environment} from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TestsRunnerService {

  constructor() { }

  public run<T>(filename: string): Promise<T> {
    return new Promise<T>((resolve, reject) => {
      const ws = new WebSocket(
        `ws://${
          window.location.host}/${environment.apiPrefix}/run-tests/?filename=${filename}`,
      );

      ws.onmessage = message => resolve(JSON.parse(message.data));

      ws.onerror = err => reject(err);
    });
  }
}
