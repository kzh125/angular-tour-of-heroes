import { Component, OnInit } from '@angular/core';
import { Observable, of, Subject } from 'rxjs';
import {
  debounceTime, distinctUntilChanged, switchMap
} from 'rxjs/operators';

import { Hero } from '../hero';
import { HeroService } from '../hero.service';

@Component({
  selector: 'app-hero-search',
  templateUrl: './hero-search.component.html',
  styleUrls: ['./hero-search.component.css']
})
export class HeroSearchComponent implements OnInit {
  heroes$: Observable<Hero[]>;
  private searchTerms = new Subject<string>();

  constructor(private heroService: HeroService) { }

  ngOnInit(): void {
    this.heroes$ = this.searchTerms.pipe(
      // wait 300ms after each keystroke before considering the term
      debounceTime(300),

      // ignore new term if same as previous term
      distinctUntilChanged(),

      // switch to new search observable each time the term changes
      switchMap((term: string) => this.heroService.searchHeroes(term)),
    );

    this.searchTerms.subscribe(
      val => console.log('s1:', val)
    );

    this.searchTerms.subscribe(
      val => ((val) => {
        return of(val);
      })(val).subscribe(val => console.log('s2: ' + val))
    );

    this.searchTerms.pipe(
      // wait 300ms after each keystroke before considering the term
      // 例如在输入框中输入a后，不到300ms就输入b
      debounceTime(300),

      // ignore new term if same as previous term
      // 例如在输入框中输入a后，选中a再输入a
      distinctUntilChanged(),

      // switch to new search observable each time the term changes
      // 例如在输入框先输入a再输入b，但是b先返回，a后返回，待a返回后舍弃a的结果
      switchMap((term: string) => {
        console.log('s3.log: ' + term);
        return of('s3.return: ' + term);
      }),
    ).subscribe(
      val => console.log('s3.result: ' + val)
    );

  }

  search(term: string) {
    this.searchTerms.next(term);
  }

}
