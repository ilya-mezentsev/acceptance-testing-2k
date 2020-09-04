import { Injectable } from '@angular/core';
import {Md5} from 'ts-md5';

@Injectable({
  providedIn: 'root'
})
export class HashService {
  public getRandomHash(): string {
    return `${Md5.hashStr((new Date()).getMilliseconds().toString())}${Math.random()}`;
  }
}
