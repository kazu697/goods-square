import { Button, Grid } from '@material-ui/core';
import React from 'react';
import ProductCard from '../components/ProductCard';

const Home: React.FC = () => {
  return (
    <Grid container spacing={2}>
      {[0, 0, 0, 0, 0, 0].map((item) => (
        <Grid item xs={12} md={6} lg={2} key={item}>
          <Button>
            <ProductCard />
          </Button>
        </Grid>
      ))}
    </Grid>
  );
};

export default Home;
