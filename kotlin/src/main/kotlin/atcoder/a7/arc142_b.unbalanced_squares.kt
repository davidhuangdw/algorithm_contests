package atcoder.a7

fun main(){
    fun input_ints(): List<Int> = readLine()!!.split(" ").map { it.toInt() }

    val (N, ) = input_ints()
    for(i in 0 until N){
        val st = i*N
        val d = (N+1)/2
        println((st+1 .. st+d).flatMap { listOf(it, it+d) }.subList(0, N).joinToString(" "))
    }
}