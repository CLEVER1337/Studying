#include <iostream>
#include <stdio.h>

#define MAX(a, b) (a > b) ? a : b

#define MIN(a, b) (a < b) ? a : b

using namespace std;
//отсортировать массив по возрастанию, если максимальный элемент встреяается раньше минимального
//матрица 5на5 определить может ли она являться магическим квадратом(сумма эл каждой строки, столбца и диагонали совпадает)

int arr_min(size_t n,  int arr[]){
	size_t m_index = 0;
	int m = arr[m_index];
	for(int i = 1; i < n; i++){
		if(arr[i] < m){
			m_index = i;
			m = arr[m_index];
		}
	}
	return  m_index;
}

int arr_max(size_t n,  int arr[]){
	size_t m_index = 0;
	int m = arr[m_index];
	for(int i = 1; i < n; i++){
		if(arr[i] > m){
			m_index = i;
			m = arr[m_index];
		}
	}
	return  m_index;
}

void quick_sort(int* array, int low, int high)
{
    int i = low;
    int j = high;
    int mid = array[(i + j) / 2];
    int temp;

    while (i <= j)
    {
        while (array[i] < mid)
            i++;
        while (array[j] > mid)
            j--;
        if (i <= j)
        {
            temp = array[i];
            array[i] = array[j];
            array[j] = temp;
            i++;
            j--;
        }
    }
    if (j > low)
        quick_sort(array, low, j);
    if (i < high)
        quick_sort(array, i, high);
}

void arr_reverse(size_t n, int* arr){
	int tmp[n];
	for(int i = 0, j = n - 1; i < n; i++, j--)
		tmp[i] = arr[j];
	for(int i = 0; i < n; i++)
		arr[i] = tmp[i];
}

int main(){
	int n = 5;
	int* a = (int*)malloc(n * sizeof(int));

	for(int i = 0; i < n; i++)
		cin >> a[i];
	
	int max_ind = arr_max(n, a);
	int min_ind = arr_min(n, a);
	quick_sort(a, 0, n - 1);
	if(max_ind > min_ind)	
		arr_reverse(n, a);

	for(int i = 0; i < n; i++)
		cout << a[i] << endl;	

    return 0;
}

