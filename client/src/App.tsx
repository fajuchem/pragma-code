import React, { useCallback, useEffect, useState } from 'react';
import { Products } from './products/Products';

export const App = () => {
    const [ loading, setLoadingState ] = useState(true);

    const setLoading = useCallback((flag: boolean) => {
        setLoadingState(flag);
    }, []);


    return (
      <>
        <header>
          <h1>SensorTech</h1>
        </header>
        <main>
          {loading &&
            <p><em>Loading...</em></p>
          }
          <Products loading={loading} setLoading={setLoading} />
        </main>
      </>
    );
}

