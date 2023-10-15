import { Component } from '@angular/core';
import { CompetitorService, Competitor } from './competitor.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'frontend';
  constructor(private competitorService: CompetitorService) {}
  showCompetitors() {
    console.log('world');
    this.competitorService
      .getCompetitors()
      .subscribe((data: any) => console.log(data));
  }
  AferViewInit() {
    this.showCompetitors();
  }
}
