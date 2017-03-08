import { Injectable } from '@angular/core';
import { Http} from '@angular/http';
import 'rxjs/add/operator/toPromise'
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import {Exam} from "../model/Exam";

@Injectable()
export class ExamListService {

  private examListUrl = 'http://localhost:8112/api/exam/list';

  constructor(private http: Http) { }

  getList(): Promise<Exam[]> {
    return this.http.get(this.examListUrl)
      .toPromise()
      .then(response => response.json() as Exam[])
      .catch(this.handleError)
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error)
  }

}
