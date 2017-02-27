import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Location } from '@angular/common';

import 'rxjs/add/operator/switchMap';

import { Exam } from '../model/exam';
import {ExamService} from "./exam.service";
import {Book} from "../model/Book";
@Component({
  selector: 'app-exam',
  templateUrl: './exam.component.html',
  styleUrls: ['./exam.component.scss']
})
export class ExamComponent implements OnInit {

  exam: Exam[] = [];
  book: Book[] = [];
  questionCount: number;

  constructor(private route: ActivatedRoute, private examService: ExamService) { }

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => {
      return this.examService.getExam(+params['examId'])
    })
      .subscribe(exam => {
        this.exam.push(exam);
        this.book.push(exam.book);
        this.questionCount = this.exam[0].question ? this.exam[0].question.length : 0;

        this.book[0].created_at_display = (new Date(this.book[0].created_at).toLocaleDateString());
        this.book[0].updated_at_display = (new Date(this.book[0].updated_at).toLocaleDateString());
        this.exam[0].created_at_display = (new Date(this.exam[0].created_at).toLocaleDateString());
        this.exam[0].updated_at_display = (new Date(this.exam[0].updated_at).toLocaleDateString());
      })
  }

}
