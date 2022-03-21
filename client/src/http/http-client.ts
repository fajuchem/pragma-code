const BASE_URL = 'http://localhost:4000/api';

export async function getAll<T>(endpoint: string): Promise<T[]> {
  try {
    const response = await fetch(`${BASE_URL}/${endpoint}`);
    return response.json();
  } catch (e) {
    console.log(`Error fetching ${endpoint}`);
    return [];
  }
}
