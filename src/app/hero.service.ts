import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Hero } from './hero';
import { HEROES } from './mock-heroes';
import { MessageService } from './message.service';

@Injectable({
  providedIn: 'root'
})
export class HeroService {
  private heroesUrl = 'http://127.0.0.1:4201/api/heroes';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient, private messageService: MessageService) { }

  getHeroes(): Observable<Hero[]> {
    this.log('fetched heroes');
    return this.http.get<Hero[]>(this.heroesUrl);
  }

  getHero(id: number): Observable<Hero> {
    this.log(`fetched hero id=${id}`);
    const url = `${this.heroesUrl}/${id}`
    return this.http.get<Hero>(url);
  }

  updateHero(hero: Hero): Observable<any> {
    this.log(`updated hero id=${hero.id}`);
    return this.http.put(this.heroesUrl, hero, this.httpOptions);
  }

  deleteHero(hero: Hero | number): Observable<any> {
    const id = typeof hero === "number" ? hero : hero.id;
    this.log(`deleted hero id=${id}`);
    const url = `${this.heroesUrl}/${id}`
    return this.http.delete(url);
  }

  addHero(hero: Hero): Observable<Hero> {
    this.log(`added hero id=${hero.id}`);
    return this.http.post<Hero>(this.heroesUrl, hero, this.httpOptions);
  }

  private log(message: string) {
    this.messageService.add(`HeroService: ${message}`);
  }
}
