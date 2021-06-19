package google_kickstart.`2020`.F

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n, x) = input_ints()
        val arr = input_ints().withIndex().map{
            val (i, v) = it
            (v+x-1)/x to i+1
        }.sortedWith(compareBy({it.first}, {it.second})).map{it.second}

        println("Case #${it}: ${arr.joinToString(" ")}")
    }
}
