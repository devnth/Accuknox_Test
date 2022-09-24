# Program to find top three foods in a restaurant

Program reads data from text files and store them in hashmap data type in key value list pairs. From that hashmap it calculates the length of the list and prints the corresponding top three food name. 
It has four functions to get the desired output.




## The four functions
-  getLogData: To read and store data from the restaurant log file in hashmap structuture type with foodmenu_id as key and eater_id as list of values. It also checks whether the diner has taken the same food twice and returns error for the same.

-  getfoodMenu: To read and store the foodmenu_id and food names from the file and store in hashmap structure with foodmenu_id as key and food name as value.

-  getTopThree: To calculate the top three selling fooditems by calculating the length of list of each keys in the hashmap of logdata.

-  printFoodName: This function prints the list of names of foods from the foodMap data in accordance with the corresponding top three ranking foodmenu_id.

## Example 1 (success)

An example of restaurant log data with no repeating pairs of foodmenu_id and eatter_id.
```bash
eater_id:    food_menu_id
--------------------------
dev12: cb100, lj20, ic60
jude23: cb100, lj20, ic60
midl26: md250, fl60
mun22: dos36, lj20
sal24: alf260, lj20
rah24: md250, lj20
gok22: cb100, lj20, ic60
fah26: alf260, ic60
deep27: cb100, lj20, alf260
anu23: alf260, ic60, fl60
```

The output:
```Food id: ic60, count: 5
Food id: md250, count: 2
Food id: fl60, count: 2
Food id: dos36, count: 1
Food id: alf260, count: 4
Food id: cb100, count: 4
Food id: lj20, count: 7


First:  Lime Juice      --- 7
Second:  Icecream       --- 5
Third:  AlFaham  Chicken Biriyani       --- 4
```

##  Example 2 (Error)

An example of restaurant log data with repeating pairs of foodmenu_id and eatter_id.
```bash
eater_id:   foodmenu_id
--------------------------
dev12: cb100, lj20, ic60
jude23: cb100, lj20, ic60
midl26: md250, fl60
mun22: dos36, lj20
sal24: alf260, lj20
rah24: md250, lj20
gok22: cb100, lj20, ic60
fah26: alf260, ic60
deep27: cb100, lj20, alf260
anu23: alf260, ic60, fl60
midl26: fl60 
```

The output:
```2022/09/23 20:13:49 Dinner eaten same food twice.
exit status 1
```
