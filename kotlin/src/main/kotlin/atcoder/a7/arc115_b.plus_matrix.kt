package atcoder.a7

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N) = inputInts()
    val C = (0 until N).map { inputInts() }

    var (mi, mj) = 0 to 0
    var mn = C[0][0]
    for(i in 0 until N)
        for(j in 0 until N){
            if(C[i][j] < mn){
                mi = i
                mj = j
                mn = C[i][j]
            }
        }

    val A = IntArray(N)
    val B = IntArray(N)
    A[mi] = 0
    B[mj] = C[mi][mj]
    for(i in 0 until N){
        A[i] = C[i][mj] - B[mj]
        B[i] = C[mi][i] - A[mi]
    }

    if((0 until N).all{i -> (0 until N).all{j -> C[i][j] == A[i] + B[j]}}){
        println("Yes")
        println(A.joinToString(" "))
        println(B.joinToString(" "))
    }else{
        println("No")
    }
}
