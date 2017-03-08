import { Injectable } from '@angular/core';
import { Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'
import {Book} from "../model/Book";
import {Exam} from "../model/Exam";
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';


@Injectable()
export class BookService {

  private bookListUrl = 'http://localhost:8112/api/book/';
  private examCreateUrl = 'http://localhost:8112/api/exam/create';

  constructor(private http: Http) {}

  getList(): Promise<Book[]> {
    return this.http.get(this.bookListUrl)
      .toPromise()
      .then(response => response.json() as Book[])
      .catch(this.handleError)
  }

  submitExamCreation(exam:Exam): Promise<Exam> {
    var examJson: any;
    examJson = {
      book_id: exam.book.id,
      name: exam.name
    };

    return this.http.post(this.examCreateUrl, JSON.stringify(examJson))
      .toPromise()
      .then(response => response.json() as Exam)
      .catch(this.handleError)
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error)
  }
}
