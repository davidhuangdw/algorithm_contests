package atcoder.a3

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_longs() = split_input().map { it.toLong() }

    fun ceil(a: Long, b: Long): Long{
        return (a+b-1)/b
    }

    val (_, L, W) = input_longs()
    val A = input_longs().toMutableList()
    A.add(L)
    var sum =  ceil(A[0], W)
    for((x, y) in A.zipWithNext()){
        val d = y - (x+W)
        if(d > 0){sum += ceil(d, W)}
    }
    println(sum)
}