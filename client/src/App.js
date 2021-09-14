import React, { Component } from 'react';

export class App extends Component {
  static displayName = App.name;

  constructor(props) {
    super(props);
    this.state = { products: [], loading: true };
  }

  componentDidMount() {
    setInterval(() => {
      this.fetchProducts();
    }, 5000);
  }

  async fetchProducts() {
    const response = await fetch('api/products');
    const data = await response.json();
    this.setState({ products: data, loading: false });
  }

  render() {
    let contents = this.state.loading
      ? (
        <p><em>Loading...</em></p>
      ) : (
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
            {this.state.products.map((product) => (
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

    return (
      <>
        <header>
          <h1>SensorTech</h1>
        </header>
        <main>
          {contents}
        </main>
      </>
    );
  }
}
