import { ApplicationConfig } from '@angular/core';
import { provideRouter, withComponentInputBinding } from '@angular/router';

import { routes } from './app.routes';
import { provideHttpClient, withJsonpSupport, withXsrfConfiguration } from '@angular/common/http';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes, withComponentInputBinding()),
    provideHttpClient()
  ]
};

export const apiEndpoint = "https://localhost:7071";      // dotnet
// export const apiEndpoint = "http://127.0.0.1:3000";    // go fiber