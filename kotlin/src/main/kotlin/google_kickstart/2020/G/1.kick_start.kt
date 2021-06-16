package google_kickstart.`2020`.G

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val KICK = "KICK"
    val START = "START"
    val T = readLine()!!.toInt()
    (1..T).forEach {
        var sum = 0
        var kicks = 0
        val str = readLine()!!
        val n = str.length
        for(i in str.indices){
            if(str.substring(i, minOf(i+4, n)) == KICK)
                kicks += 1
            else if(str.substring(i, minOf(i+5, n)) == START)
                sum += kicks
        }
        println("Case #${it}: $sum")
    }
}
