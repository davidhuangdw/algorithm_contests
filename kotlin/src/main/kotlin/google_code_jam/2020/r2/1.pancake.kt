package google_code_jam.`2020`.r2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        var (L, R) = input_longs()
        var i = 1
        while(true){
            if(i > maxOf(L, R)) break
            if(L >= R) L -= i
            else R-= i
            i += 1
        }
        println("Case #${it}: ${i-1} $L $R")
    }
}
