import { Injectable } from '@angular/core';
import {environment} from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TestsRunnerService {
  private readonly maxFileSize = 32 * 1024 * 1024;  // 32 MB
  private readonly wsProtocol = window.location.protocol === 'https:' ? 'wss' : 'ws';

  public async getFileCheckError(file: File): Promise<string | undefined> {
    if (file.size > this.maxFileSize) {
      return 'File is too large';
    }

    const content = await file.text();
    const meaningLinesCount = content
      .split('\n')
      .filter(s => !!s.trim())
      .length - 1;

    if (
      environment.shouldCheckLinesCount &&
      meaningLinesCount > environment.maxLinesCount
    ) {
      return 'Maximum tests lines count exceeded';
    } else {
      return undefined;
    }
  }

  public run<T>(filename: string): Promise<T> {
    return new Promise<T>((resolve, reject) => {
      const ws = new WebSocket(
        `${this.wsProtocol}://${window.location.host}/${environment.apiPrefix}/run-tests/?filename=${filename}`,
      );

      ws.onmessage = message => resolve(JSON.parse(message.data));

      ws.onerror = err => reject(err);
    });
  }
}
