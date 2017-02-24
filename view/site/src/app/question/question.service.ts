import { Injectable } from '@angular/core';
import { Headers, Http} from '@angular/http'
import {Question} from "../model/Question";


@Injectable()
export class QuestionService {

  private questionListUrl = 'http://localhost:8112/api/book';


  constructor(private http: Http) { }

  getList(bookId: number): Promise<Question[]> {
    return this.http.get(`${this.questionListUrl}/${bookId}/questionList`)
      .toPromise()
      .then(response => response.json() as Question[])
      .catch(this.handleError)
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error)
  }

}
