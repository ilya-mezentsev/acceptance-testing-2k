import { Injectable } from '@angular/core';
import {ErrorResponse, Fetcher, FileSender, Response, ServerResponse} from '../../interfaces/fetcher';
import {environment} from '../../../environments/environment';
import {Md5} from 'ts-md5';
import {SessionStorageService} from '../session/session-storage.service';

@Injectable({
  providedIn: 'root'
})
export class FetcherService implements Fetcher, FileSender {
  private publicKey = '';
  private readonly apiPath: string = '/api/web-app';

  constructor(
    private readonly sessionStorage: SessionStorageService,
  ) {}

  private static trimSlashes(endpoint: string): string {
    if (endpoint.startsWith('/')) {
      endpoint = endpoint.substr(1);
    }

    if (endpoint.endsWith('/')) {
      endpoint = endpoint.substring(0, endpoint.length - 1);
    }

    return endpoint;
  }

  public get(endpoint: string): Promise<ServerResponse> {
    return this.fetch(
      endpoint,
      {
        method: 'GET'
      }
    );
  }

  public post(endpoint: string, data: any): Promise<ServerResponse> {
    return this.fetch(
      endpoint,
      {
        method: 'POST',
        body: JSON.stringify(data)
      }
    );
  }

  public patch(endpoint: string, data: any): Promise<ServerResponse> {
    return this.fetch(
      endpoint,
      {
        method: 'PATCH',
        body: JSON.stringify(data),
      }
    );
  }

  public delete(endpoint: string): Promise<ServerResponse> {
    return this.fetch(
      endpoint,
      {
        method: 'DELETE'
      }
    );
  }

  public sendFile<T>(endpoint: string, file: FormData): Promise<ErrorResponse | Response<T>> {
    return this.fetch(endpoint, {
      method: 'POST',
      body: file
    });
  }

  private async fetch(endpoint: string, settings: RequestInit): Promise<ServerResponse> {
    settings.headers = {
      ...settings.headers,
      'X-CSRF-Token': this.getCSRFToken(),
      'AAT-Account-Hash': this.sessionStorage.getSessionId(),
    };
    const response = await fetch(
      `${this.apiPath}/${FetcherService.trimSlashes(endpoint)}/`,
      settings
    );
    this.publicKey = response.headers.get('X-CSRF-Public-Token');

    return response.json();
  }

  private getCSRFToken(): string {
    return btoa(
      `${Md5.hashStr(environment.csrfPrivateKey)}|${Md5.hashStr(this.publicKey)}`
    );
  }
}
