---
layout: page
nav_order: 11
---

# Testing

## Unit tests

Unit tests can be started from command line by the following command:

```
make test
```

Code coverage can be generated by:

```
make cover
```

for HTML output or:

```
make coverage
```

for plain text output.

## Profiling

Profiler can be started from command line by the following command;

```
make profile
```

## BDD tests

Behaviour tests for this service are included in [Insights Behavioral
Spec](https://github.com/RedHatInsights/insights-behavioral-spec) repository.
In order to run these tests, the following steps need to be made:

1. clone the [Insights Behavioral Spec](https://github.com/RedHatInsights/insights-behavioral-spec) repository
1. go into the cloned subdirectory `insights-behavioral-spec`
1. run the `notification_writer_tests.sh` from this subdirectory

List of all test scenarios prepared for this service is available at
<https://github.com/RedHatInsights/insights-behavioral-spec#ccx-notification-service>
