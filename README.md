[![Go Reference](https://pkg.go.dev/badge/github.com/victorguarana/vehicle-routing.svg)](https://pkg.go.dev/github.com/victorguarana/vehicle-routing)
[![Go Report Card](https://goreportcard.com/badge/victorguarana/vehicle-routing)](https://goreportcard.com/report/victorguarana/vehicle-routing)
[![License](https://img.shields.io/github/license/victorguarana/vehicle-routing)](https://github.com/victorguarana/vehicle-routing/blob/main/LICENSE)

# Vehicle Routing

## About the Project

This project implements and compares multiple greedy algorithms and metaheuristics* algorithms to solve the Hybrid Vehicle Routing Problem with Trucks and Drones (HVRP-TD), a variant of the Vehicle Routing Problem (VRP), using Golang.

> Metaheuristics are high-level procedures designed to generate or select heuristics that provide sufficiently good solutions to optimization problems, especially with incomplete or imperfect information or limited computation capacity.

## Problem Setup

- **Truck**: One or more trucks with variable capacity and speed, moves by Manhattan distance;
- **Drone**: Zero or more drones with fixed capacity and speed, moves by Euclidean distance;
- **Customer**: A list of customers with demand and location;
- **Warehouse**: A list of warehouses with unlimited supply of goods;
- **Route**: One per Truck.

## Implemented Algorithms

- **Closest Neighbor**: Createa a route by moving the vehicle to closest customer based on distance to current location;
- **Best Insertion**: Creates a route by iterating through the customer list and inserting it into the best position based on total distance;
- **Drone Strike Insertion**: Improve a truck only route by replacing truck deliveries with drone deliveries
- **Greedy Coverage by Drones**: Create a route by moving the truck to the customer with more near by customers, then send the drone(s) to the nearest customer(s).
- **Heuristic based on Iterated Local Search (ILS)**: Impreve a route by applying all possible moves and selecting the best one;

## Comparison Criteria

- **Total Distante**: Calculate the total distance of the route;
- **Total Time**: Calculate the total time of the route, including the time when vehicles are stopped waiting for others;
- **Total Fuel**: Calculate the total fuel consumption of the route.

## Example
- [Example directory](./example/) contains a sample input file and multiple output file with implemented algorithms.
