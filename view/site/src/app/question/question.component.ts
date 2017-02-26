import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Location } from '@angular/common';

import {Question} from "../model/Question";
import {QuestionService} from "./question.service";

import {SelectItem} from 'primeng/primeng'

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css']
})
export class QuestionComponent implements OnInit {
  error: string;
  questions: Question[];
  bookName: string;
  questionType: SelectItem[];
  questionDifficulty: SelectItem[];
  questionSelection: Question;
  displayDetail: boolean;

  constructor(
    private questionService: QuestionService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => {
        this.bookName = params['bookName'];
        return this.questionService.getList(+params['bookId'])
      })
      .subscribe(list => {
        this.questions = list;
        for(let question of this.questions) {
          question.created_at_display = (new Date(question.created_at).toLocaleDateString());
          question.updated_at_display = (new Date(question.updated_at).toLocaleDateString());
        }
      });

    this.questionType = [];
    this.questionType.push({label: 'All', value: null});
    this.questionType.push({label: 'T/F', value: 'T/F'});
    this.questionType.push({label: 'MC', value: 'MC'});
    this.questionType.push({label: 'SA', value: 'SA'});

    this.questionDifficulty = [];
    this.questionDifficulty.push({label: 'All', value: null});
    this.questionDifficulty.push({label: 'EASY', value: 'EASY'});
    this.questionDifficulty.push({label: 'MEDIUM', value: 'MEDIUM'});
    this.questionDifficulty.push({label: 'HARD', value: 'HARD'});

  }

  showText(id: number): void {
    this.questions[id].isShow = true;
  }

  paginate(event) {
    console.error(event);
  }

  onRowSelect(event) {
    console.error(this.questionSelection);

    this.displayDetail = true;
  }

  goBack(): void {
    this.location.back();
  }


}
