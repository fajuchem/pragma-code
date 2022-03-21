import { getProducts } from "./repository";

window.fetch = jest.fn();

describe('getProducts', () => {
  it('calls right endpoint for products', async () => {
    await getProducts();
    expect(fetch).toBeCalledWith('http://localhost:4000/api/products');
  });
});
