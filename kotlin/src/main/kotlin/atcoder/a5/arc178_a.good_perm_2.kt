package atcoder.a5

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (N, ) = input_ints()
    val A = input_ints().sorted()
    if(A[0] == 1 || A.last() == N){
        println(-1)
        return
    }
    val res = (1..N).toMutableList()
//    var t = -1
//    for(i in A.size-1 downTo 0){
//        if(i+1 >= A.size || A[i+1] != A[i]+1){
//            t = A[i]+1
//        }
//        var a = A[i]
//        res[a-1] = res[t-1].also{res[t-1] = res[a-1]}
//    }
    for(a in A){
        res[a-1] = res[a].also{res[a] = res[a-1]}
    }
    println(res.joinToString(" "))
}