package atcoder.a4

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }

    val (N, ) = input_ints()
    val R = input_ints()
    val C = input_ints()

    val (q, ) = input_ints()
    val res = (1..q).map{
        val (i, j) = input_ints()
        if(C[j-1] >= N+1 - R[i-1]) '#' else '.'
    }
    println(res.joinToString(""))
}