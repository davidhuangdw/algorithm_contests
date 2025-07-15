package atcoder.a6

fun main(){
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    readLine()
    var last = 0L
    for(v in inputInts()){
        val bit = 1L shl v
        val up = bit shl 1
        var cur = (last - (last and (up-1))) or bit
        last = if(cur > last) cur else cur + up
//        println(last)
    }
    println(last)
}
