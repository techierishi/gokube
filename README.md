# orchestrator-in-go
Repository for working through [Build An Orchestrator in Go](https://www.manning.com/books/build-an-orchestrator-in-go)

## Prerequisite
* Go (v1.20)
* Linux
* Docker

## Libs
* chi (v5.0.3)
* Docker SDK (v20.10.7+incompatible)
* BoltDB (v1.3.1)
* goprocinfo


## Prepare environment

```bash
make network # only once

make build

make run

```

## How to run

```bash

make exec # To start the docker container where we can run all components

# Now we are in docker container
tmux # To start tmux session
make panes # To create 3 tmux panes each for general use, manager, worker 

# Start worker node in tmux pane on bottom right
./gokube worker

# Press ctrl+b and release, Press up arrow
# Start control plane in tmux pane on top right
./gokube manager


# Press ctrl+b and release, Press left arrow
# Run a pod which will be schedule on worker in general tmux pane on left
./gokube run hack/task.json


# To check the pod list, similar to kubectl get pods
./gokube status

```

![Screenshot](./hack/screenshot.png)

## Definitions 
* Orchestrator
  * Automates deploying, scaling & managing containers.
* Task
  * Smallest unit of work. 
    * Should specify it's needs for memory, CPU & disk. 
    * What to do in case of failure (restart_policy). 
    * Name of container image used to run the task.
* Job (not included in this implementation)
  * Group of tasks to perform a set of functions. 
  * Should specify details at a higher level than tasks. 
    * Each task that makes up the job. 
    * Which data centers the job should run in. 
    * The type of job. 
* Scheduler
  * Decides which machine can best host the tasks defined. 
  * Could be a process within the machine or possibly a different service.
  * Decisions can vary in model complexity from round-robin to score calculation based on variables and deciding node by best score 
    * Detrmine set of candidate machines to run task
    * Score machines from best to worst
    * Pick machine with best score
* Manager 
  * Brain of orchestrator & main entry point
  * Jobs go to manager, manager uses scheduler to determine where to run job. 
  * Can gather metrics, to use for scheduling.
    * Accept requests from users to start/stop tasks
    * Schedule tasks onto worker machines
    * Keep track of tasks, states & the machine they run on 
* Worker
  * Runs the tasks assigned to it, attempts task restarts on failure, makes metrics about task and machine health for master
    * Runs asks as containers
    * Accepts tasks to run from manager
    * provides relevant statistics to the manager to schedule tasks
    * Keeps track of tasks & their state 
* Cluster
  * Logical grouping of all above components. 
* CLI 
  * Main user interface
    * Start/Stop Tasks
    * Get status of tasks
    * See state of machines
    * Start manager
    * Start worker 
