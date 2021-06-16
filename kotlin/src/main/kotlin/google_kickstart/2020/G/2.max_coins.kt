package google_kickstart.`2020`.G


fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n,) = input_ints()
        val sum = MutableList(n*2){0L}
        for(i in 0 until n){
            val coin = input_longs()
            for(j in 0 until n)
                sum[n+i-j] += coin[j]
        }
        println("Case #${it}: ${sum.maxOrNull()}")
    }
}
