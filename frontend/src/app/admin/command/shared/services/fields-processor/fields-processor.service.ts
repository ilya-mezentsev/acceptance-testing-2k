import { Injectable } from '@angular/core';
import {TestCommandSettings} from '../../../../types/types';

@Injectable({
  providedIn: 'root'
})
export class FieldsProcessorService {
  private static removeTrailingSlashes(s: string): string {
    while (s.endsWith('/')) {
      s = s.substring(0, s.length - 1);
    }

    return s;
  }

  private static removeLeadingSlashes(s: string): string {
    while (s.startsWith('/')) {
      s = s.substring(1);
    }

    return s;
  }

  public prepareSettings(settings: TestCommandSettings): TestCommandSettings {
    return {
      ...settings,
      base_url: FieldsProcessorService.removeTrailingSlashes(settings.base_url),
      endpoint: FieldsProcessorService.removeLeadingSlashes(settings.endpoint),
    };
  }
}
