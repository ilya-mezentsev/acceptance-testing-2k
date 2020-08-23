import { Injectable } from '@angular/core';

type Session = {account_hash: string};

@Injectable({
  providedIn: 'root'
})
export class SessionStorageService {
  private session: Session;

  public saveSession(session: Session): void {
    this.session = session;
  }

  public getSessionId(): string {
    return this.session?.account_hash;
  }
}
