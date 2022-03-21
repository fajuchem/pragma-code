import React from 'react';
import { render, waitFor, screen } from '@testing-library/react';
import '@testing-library/jest-dom';

import { App } from './App';

jest.mock('./products/Products', () => ({
  ...jest.requireActual('./products/Products'),
  Products: ({ setLoading }: { setLoading: any }) => {
    setLoading(true);
    return (<div>products</div>);
  }
}));

describe('App', () => {
  it('render correct', async () => {
    render(<App />);

    await waitFor(() => expect(screen.getByText('SensorTech')).toBeInTheDocument());
  });

  it('show loading', async () => {
    render(<App />);

    await waitFor(() => expect(screen.getByText('Loading...')).toBeInTheDocument());
  });

  it('render products component', async () => {
    render(<App />);

    await waitFor(() => expect(screen.getByText('products')).toBeInTheDocument());
  });
});
