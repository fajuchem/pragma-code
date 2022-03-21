import { getAll } from './http-client';

window.fetch = jest.fn();

describe('get method', () => {
  beforeEach(() => {
    (fetch as jest.Mock).mockClear();
  });

  it('return data correctly', async () => {
    (fetch as jest.Mock).mockResolvedValue({
      json: () =>
        Promise.resolve([
          {
            id: '1',
            name: 'Pilsner',
            minimumTemperature: 4,
            maximumTemperature: 6,
            temperature: 1,
            temperatureStatus: 'too low',
          },
          {
            id: '2',
            name: 'IPA',
            minimumTemperature: 5,
            maximumTemperature: 6,
            temperature: -2,
            temperatureStatus: 'too low',
          },
        ]),
    });
    const data = await getAll('endpoint');
    expect(data).toHaveLength(2);
    expect(data[0]).toEqual({
      id: '1',
      name: 'Pilsner',
      minimumTemperature: 4,
      maximumTemperature: 6,
      temperature: 1,
      temperatureStatus: 'too low',
    });
  });

  it('calls correctly endpoint', async () => {
    (fetch as jest.Mock).mockResolvedValue({ json: () => {} });
    await getAll('endpoint');
    expect(fetch).toBeCalledWith('http://localhost:4000/api/endpoint');
  });

  it('fetch fails', async () => {
    const data = await getAll('endpoint');
    expect(data).toEqual([]);
  });
});
