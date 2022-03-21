import React from 'react';
import { render, waitFor, screen } from '@testing-library/react';
import '@testing-library/jest-dom';

import { getProducts } from './repository';
import { Products } from './Products';

jest.mock('./repository');

describe('Products', () => {
  const loading = true;
  const setLoading = jest.fn();
  const defaultProps = {
    loading,
    setLoading,
  }

  it('render correct', async () => {
    (getProducts as jest.Mock).mockResolvedValue([{ id: 1, name: 'fe', temperature: 1, temperatureStatus: 'good' }]);
    render(<Products { ...defaultProps} loading={false} />);

    await waitFor(() => expect(screen.getByText('good')).toBeInTheDocument());
  });

  it('setLoading when fetching data', async () => {
    (getProducts as jest.Mock).mockResolvedValue([{ id: 1, name: 'fe', temperature: 1, temperatureStatus: 'good' }]);
    render(<Products { ...defaultProps} loading={true} />);

    await waitFor(() => expect(setLoading).toBeCalledWith(false));
  });
});
