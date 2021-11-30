import { Pipe, PipeTransform } from '@angular/core';

import { Product } from '../models/products.model';

@Pipe({
  name: 'sort',
  pure: true
})
export class SortPipe implements PipeTransform {

  transform(products: Product[], sort: string): Product[] {
    if (sort === 'asc' || sort === 'desc') {
      products.sort((a, b) => {
        const diff = a.price - b.price;

        if (diff === 0) {
          return 0
        }

        const sign = Math.abs(diff) / diff;

        return sort === 'asc' ? sign : -sign;
      });
    }
    return products;
  }

}
