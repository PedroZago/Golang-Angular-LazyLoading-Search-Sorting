import { Component, Input, EventEmitter, Output } from '@angular/core';

import { Filter } from '../models/filter.model';
import { Product } from '../models/products.model';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent {

  @Input('products') products: Product[] = [];
  @Input('filter') filter!: Filter;
  @Input('lastPage') lastPage!: number;

  @Output('setFilters') setFilters = new EventEmitter();

  search(event: any): void {
    const s = event.target.value;
    this.setFilters.emit({
      ...this.filter,
      s,
      page: 1
    });
  }

  sort(event: any): void {
    const sort = event.target.value;
    this.setFilters.emit({
      ...this.filter,
      sort,
      page: 1
    });
  }

  loadMore(): void {
    this.setFilters.emit({
      ...this.filter,
      page: this.filter.page + 1
    });
  }

}
