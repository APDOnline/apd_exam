import { Injectable } from '@angular/core';
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'
import {Book} from "../model/Book";

@Injectable()
export class BookService {

  private bookListUrl = 'http://localhost:8112/api/book/';

  constructor(private http: Http) { }

  getList(): Promise<Book[]> {
    return this.http.get(this.bookListUrl)
      .toPromise()
      .then(response => response.json() as Book[])
      .catch(this.handleError)
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error)
  }
}
