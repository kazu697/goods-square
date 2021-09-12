import { Button, Grid } from '@material-ui/core';
import React from 'react';
import { makeStyles } from '@material-ui/core';
import { useHistory } from 'react-router';
import ProductCard from '../components/ProductCard';

const useStyles = makeStyles((theme) => ({
  root: {
    [theme.breakpoints.up('md')]: {
      margin: '0 10%',
    },
  },
}));

const Home: React.FC = () => {
  const history = useHistory();
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <h1>新着</h1>
      <Grid container spacing={2}>
        {[0, 0, 0, 0, 0, 0].map((item) => (
          <Grid item xs={12} md={6} lg={3} key={item}>
            <Button onClick={() => history.push(`/product/${0}`)}>
              <ProductCard />
            </Button>
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export default Home;
