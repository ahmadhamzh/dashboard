import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { tokenNotExpired } from 'angular2-jwt';

@Injectable()
export class Auth {
  constructor() {
    const token = this.getTokenFromQuery();
    if (token) {
      localStorage.setItem('token', token);
    }
  }

  public getBearerToken(): string {
    return localStorage.getItem('token');
  }

  public authenticated() {
    // Check if there's an unexpired JWT
    // This searches for an item in localStorage with key == 'token'
    return tokenNotExpired('token');
  }

  public logout() {
    localStorage.removeItem('token');
  }

  private getTokenFromQuery(): string {
    const results = new RegExp('[\?&]id_token=([^&#]*)').exec(window.location.href);
    return results == null ? null : results[1] || '';
  }
}
