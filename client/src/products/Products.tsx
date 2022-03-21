import React, { useEffect, useState } from 'react';
import { getProducts, Product } from './repository';

export type ProductsProps = {
  loading: boolean;
  setLoading: Function;
}

export const Products = (props: ProductsProps) => {
  const { loading, setLoading } = props;
  const [ products, setProducts ] = useState<Product[]>([]);
  const update = () => getProducts().then(data => {
    setProducts(data);
    setLoading(false);
    setTimeout(update, 2000);
  });

  useEffect(() => {
    update();
  }, []);

  if (loading) {
    return null;
  }

  return (
    <>
      <h2>Beers</h2>
      <table>
        <thead>
          <tr>
            <th align="left">Product</th>
            <th align="left">Temperature</th>
            <th align="left">Status</th>
          </tr>
        </thead>
        <tbody>
          {products.map((product: Product) => (
            <tr key={product.id}>
              <td width={150}>{product.name}</td>
              <td width={150}>{product.temperature}</td>
              <td width={150}>{product.temperatureStatus}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}

