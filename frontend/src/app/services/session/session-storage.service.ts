import { Injectable } from '@angular/core';

type Session = {account_hash: string};

@Injectable({
  providedIn: 'root'
})
export class SessionStorageService {
  private session: Session;

  public deleteSession(): void {
    this.saveSession(null);
  }

  public hasSession(): boolean {
    return !!this.session;
  }

  public saveSession(session: Session): void {
    this.session = session;
  }

  public getSessionId(): string {
    return this.session?.account_hash;
  }
}
