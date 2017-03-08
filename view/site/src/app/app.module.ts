import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {HttpModule} from '@angular/http';
import {
  ButtonModule,
  DataTableModule,
  SharedModule,
  TabMenuModule,
  TabViewModule,
  DropdownModule,
  DialogModule,
  InputTextModule,
  GrowlModule
} from 'primeng/primeng';

import {AppComponent} from './app.component';
import {BookComponent} from './book/book.component';
import {RouteModule} from "./route/route.module";
import {BookService} from "./book/book.service";
import {QuestionComponent} from './question/question.component';
import {QuestionService} from './question/question.service';
import {ExamComponent} from './exam/exam.component';
import {requestOptionsProvider} from './default-request-options.service';
import {ExamService} from './exam/exam.service';
import { ExamQuestionComponent } from './exam-question/exam-question.component';
import { ExamListComponent } from './exam-list/exam-list.component';
import {ExamListService} from "./exam-list/exam-list.service";

@NgModule({
  declarations: [
    AppComponent,
    BookComponent,
    QuestionComponent,
    ExamComponent,
    ExamQuestionComponent,
    ExamListComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouteModule,
    DataTableModule,
    SharedModule,
    ButtonModule,
    TabMenuModule,
    DropdownModule,
    DialogModule,
    InputTextModule,
    GrowlModule,
    TabViewModule
  ],
  providers: [
    BookService,
    QuestionService,
    ExamService,
    requestOptionsProvider,
    ExamListService

  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
