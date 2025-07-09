package atcoder.a4

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (N, ) = input_ints()
    val (A, B, C) = (1..3).map{input_ints().sorted()}
    var (j, k) = 0 to 0
    for((i, a) in A.withIndex()){
        j = (j until N).find{B[it]>a} ?: N
        if(j >= N){
            println(i)
            return
        }
        k = (k until N).find{C[it]>B[j]} ?: N
        if(k >= N){
            println(i)
            return
        }
        j++
        k++
    }
    println(N)
}