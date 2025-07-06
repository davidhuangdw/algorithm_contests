package atcoder.a3

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}

    val (N, ) = input_ints();
    val A = input_ints().sorted();
    val res = (0 until N).map{ A[(if(it and 1 == 0) 0 else (N+1)/2) +it/2]}
    val poss = (1 until N step 2).all{i -> res[i-1]<res[i] && res[i] > res[i+1]}
    println(if(poss) "Yes" else "No")
}