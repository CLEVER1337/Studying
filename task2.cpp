#include <iostream>

using namespace std;

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

int main(){
	int n = 5;

	int arr[n][n];

	for(int i = 0; i < n; i++)
		for(int j = 0; j < n; j++)
			cin >> arr[i][j];
	
	int sums[2*n];
	for(int i = 0; i < 2*n; i++)
		sums[i] = 0;

	int diag = 0, diag1 = 0;
	for(int i = 0; i < n; i++){
		for(int j = 0; j < n; j++){
			if(i == j)
				diag += arr[i][j];
			if((i + j) == n - 1)
				diag1 += arr[i][j];
			sums[i] += arr[i][j];
			sums[n + i] += arr[j][i];
		}
	}

	if(diag == diag1)
		if(arr_min(2*n, sums) == arr_max(2*n, sums))
			cout << "MAGIC" << endl;
}

