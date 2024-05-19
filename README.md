# Test task for YADRO

--------------
## For run app:
1) Clone repository  
``git clone https://github.com/sabina301/Test-Task-1``  
  

2) Go to root directory  
``cd Test-Task-1``


3) Go to tz-yadro directory  
   ``cd tz-yadro``


4) Compile the executable file from the root directory  
``go build ./cmd/``  


5) Run executable file   
``./cmd``  

## For run tests:

1) Go to internal from the tz-yadro directory  
``cd internal``  


2) Run tests  
``go test``
  
  
___

### Алгоритм решения поставленной задачи:
1) Считаем количество шаров каждого цвета и количество шаров, находящихся в каждом контейнере. Таким образом заполняем слайсы.
2) Сортируем и сравниваем получившиеся значения в слайсах. Если идентичны, то все верно. Иначе неверно.
3) По результатам предыдущего действия печатаем "yes" или "no" (если верно "yes", иначе "no").

### Оценка сложности алгоритма:
`O(n^2 + nlogn)`, где n - количество контейнеров.
