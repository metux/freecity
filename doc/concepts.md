
power grid
===========

* we can have multiple separate power grids
* each one can have multiple consumers and suppliers
* list of connections, eg. buildings or powerlines
* resistance / voltage drop within a network is ignored (realistic since distances in a city are short)
* only one power level: low voltage grid (220/110), no simulation of middle and high voltage
  (possibly in later versions: middle grid between cities - will act as producers)
* on overload: load shedding on individual consumers

buildings
=========

* variable list of buildings of various types
* position is top-left edge

aging and modernization
-----------------------

* some buildings mage age and finally collapse when maximium lifetime reached
* those can be modernized, which will reset the age
    * standard ruleset: costs: build-price / lifetime * (lifetime - age)
* those with maintenance costs are automatically modernized (as long as the costs are paid)

revenue
-------

* some buildings (eg. residential, industrial, ...) may give revenue
* these have negative costs in the ruleset
* actual revenue is calculated by tax rates (which in turn influences demand)
