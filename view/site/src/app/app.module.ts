import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {HttpModule} from '@angular/http';
import {DataTableModule,SharedModule} from 'primeng/primeng';
import {ButtonModule} from 'primeng/primeng';

import {AppComponent} from './app.component';
import {BookComponent} from './book/book.component';
import {RouteModule} from "./route/route.module";
import {BookService} from "./book/book.service";
import {QuestionComponent} from './question/question.component';
import {QuestionService} from './question/question.service';

@NgModule({
  declarations: [
    AppComponent,
    BookComponent,
    QuestionComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouteModule,
    DataTableModule,
    SharedModule,
    ButtonModule
  ],
  providers: [
    BookService,
    QuestionService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
