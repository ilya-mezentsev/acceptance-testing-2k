import { Injectable } from '@angular/core';
import {TestCommandSettings} from '../../../../types/types';

@Injectable({
  providedIn: 'root'
})
export class FieldsProcessorService {
  private static removeTrailingSlash(s: string): string {
    return s.endsWith('/')
      ? s.substring(0, s.length - 1)
      : s;
  }

  private static removeLeadingSlash(s: string): string {
    return s.startsWith('/')
      ? s.substring(1)
      : s;
  }

  public prepareSettings(settings: TestCommandSettings): TestCommandSettings {
    return {
      ...settings,
      base_url: FieldsProcessorService.removeTrailingSlash(settings.base_url),
      endpoint: FieldsProcessorService.removeLeadingSlash(settings.endpoint),
    };
  }
}
