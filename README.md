# Use continuous profiling to gain a deeper understanding of your incidents

This is a demo I presented at SRE Day London 15th September 2023. Slides can be found [here](https://docs.google.com/presentation/d/1DRr7AmOMPch-Zh6hVpj70EOG9sHlrmVoyaTpXk3QTfs/edit?usp=sharing)

The go application is adapted from https://github.com/felixge/fgprof. Thanks for his fantastic work in the profiling/performance space.

## Usage

### Spin up environment

```
# Connect/Create to a kubernetes test cluster
# e.g. $ kind cluster create

# Install pyroscope via helm chart
$ helm upgrade --install pyroscope grafana/pyroscope

# Install demo app
$ kubectl apply -f deployment/

# Now connect to the pyroscope UI
$ kubectl port-forward pyroscope-0 4040
```
