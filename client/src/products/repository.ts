import { getAll } from '../http/http-client';

export interface Product {
  id: string;
  name: string;
  minimumTemperature: number,
  maximumTemperature: number,
  temperature: number,
  temperatureStatus: string
}

export function getProducts() {
  return getAll<Product>('products');
}
