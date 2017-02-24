import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Location } from '@angular/common';

import {Question} from "../model/Question";
import {QuestionService} from "./question.service";

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css']
})
export class QuestionComponent implements OnInit {
  error: string;
  questions: Question[];

  constructor(
    private questionService: QuestionService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => this.questionService.getList(+params['bookId']))
      .subscribe(list => this.questions = list )
  }

  goBack(): void {
    this.location.back();
  }

}
