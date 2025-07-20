package atcoder.a8

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N) = inputInts()
    val A = IntArray(N+1)
    for(i in 1..N){
        var cnt = 1
        var r = i
        for(p in 2 until i){
            if(p * p > r){
                break
            }
            while(r %p == 0){
                r /= p
                cnt ++
            }
        }
        if(r > 1){cnt ++}
        A[i] = cnt
    }

    println(A.slice(1..N).joinToString(" "))
}
